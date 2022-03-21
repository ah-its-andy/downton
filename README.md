# downton
## Roadmap
- Collections : Generic collection types 
  - Type supplies:
    - Iterator
      - Cast collection for array/slice traversal
    - List (interface)
      - ArrayList (List's implemention)
        - Features:
          - Add an element to the end of collection
          - Remove an element with element's index
          - Find index of an element
          - Find an element with index
        - APIs:
          - Add
          - Remove
          - RemoveAt
          - IndexOf
          - Clear
          - ToArray
          - Count
          - Get
          - GetCapacity          
      - LinkedList (List's implemention)
        - Features: 
          - Add an element to the front of collection
          - Add an element to the end of collection
          - Insert an element before another element
          - Insert an element after another element
          - Remove an element with index
          - Find index of an element
          - Find an element with index
          - Chained list style api
        - APIs:
          - 
      - SkipList (List's implemention) 
    - Map (interface)
      - HashMap (Map implemention with out go map)
      - OrderedMap (Map implemention with go map)
    - Set (interface)
      - HashSet (Set's implemention)
    - Querable
      - Chained-Api with Linq style
  
- Configuration : Application configs lib
  - Features:
    - File type supportives: xml,json,yaml
    - Get value with path
    - Get value with chained-api
    - Cast config to a struct type (include struct tags)

- ServiceContext : Dependency Injection Extensions
  - Features:
    - Lifetime supportives: singleton, scope, transient
    - Cycle references detection
    - Struct field injection tag
    - Custom service initialization function 
  
- Hosting : Abstracts hosts definitions(such as http service, console application)
  - HostBuilder
  - Global panic&recover handler
  - Dependency Injection Supportive with ServiceContext

- Web : Http server 
  - WebHost : Hosting implemention
  - Middlewares : Http middleware 
  - Request parameters auto-binding

