# ドメインモデル図

```mermaid
classDiagram

namespace EventAggregation {
  class Event {
    UUID id
    string title
    string description
    Venue venue
    datetime start_time
    datetime end_time
  }

  class Venue {
    UUID id
    string name
    string address
    int capacity
  }
}

namespace UserAggregation {
	class User {
    UUID id
    string name
    string email
    Role role
  }
  
  class Role {
    <<enumeration>>
    Organizer
    Attendee
  }
}

namespace RegistrationAggregation {
	class Registration {
    UUID id
    Status status
    datetime registered_at
    Ticket ticket
  }

  class Ticket {
    UUID id
    string qr_code
    TicketStatus status
    datetime issued_at
    datetime used_at
  }

  class Status {
    <<enumeration>>
    Registered
    Canceled
  }

  class TicketStatus {
    <<enumeration>>
    Active
    Used
    Canceled
  }
}

%% Relationships
Event --* Venue : Event作成時にVenueも存在しないといけない
Event --> Registration : capacityを超えた登録はできない
Event --> User : Organizerロールを持つUserがEventを作成できる
Registration --> User
Registration --> Ticket : 登録時にTicketも存在しないといけない<br>登録がキャンセルされた時Ticketもキャンセルする必要がある
User --> Role
Registration --> Status
Ticket --> TicketStatus
```

# 集約とは何か

集約を簡単に表すと**リポジトリへの入出力単位**のこと。

つまり集約ないのエンティティは常に一緒に作成、更新が行われる。

なのでエンティティ間に確保したい整合性がある場合、それらを集約として一つのまとまりとする。

# なぜ集約を使うのか？

結論、実装の手間を省きコードを簡潔にするため。

上記のドメインモデル図を例に仮にエンティティごとをそれぞれ集約として仮定した場合、

UseCase層が複雑になり、実装コストが上がる。

<実装コストが上がる例を上げる>

ただし、多くのエンティティを集約でまとめてしまうと以下のような弊害が生まれる。

<弊害が生まれる例を記載>

なので集約に複数のエンティティを含む場合は、上記の弊害を考慮した上で決める。

ただし、実装してみた結果集約単位を再検討することは結構起こる。

その場合は再度ドメインモデル図に立ち返り、見直すという作業をするのが一般的。
