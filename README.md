# hexagonal

## domain
- domain
  - struct 
  - constant
- usecase
  - only about domain itself
  - dont pass other domain to be parameter on the use case
  - no repository call
- repository
  - input struct for support repository
  - interface repository

## usecase
- no public func except the main use case


## usecase/business
- define the use case object with object domain repository that needed
- can pass the domain object from parameter to the use case
- make sure the use case is about the business


Main Changelog
- breakdown the bloated responsibility on the usecase
  - domain
  - common
- streamline / remove builder

Changelog
- Interfacing 
  - usecase layer
  - common usecase layer
  - domain logic ???
- Domain Layer
  - Redesign domain layer
    - design / create a new domain
    - Use inheritance domain
    - Add method to help streamline usecase layer
      - the method should about the domain entity itself
  - Add more responsible on domain layer
    - domain -> defines the entity, const and related types.
    - usecase -> contains usecase / business about the domain only.
    - repository -> Interface and all related provider / repository 
  - Categorizing / grouping domain
    - aggregates
      - all domain that contains embed for domain
    - entities
- usecase layer
  - streamline logic by wraping all logic non usecase to func and move it other package like (ctx, monitoring, etc )
  - reordering flow ( validate first instead getting data )
  - no builder logic on usecase
  - shared repository / single repository ???
- usecase common layer
  - Categorizing by ??? by identify and make sure for single responsibility
  - shared repository / single repository ???
- others
  - Comment
  - Public / private