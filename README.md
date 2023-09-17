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
