# Custody

This is the process of taking a person into, and holding them in, custody. 

The welfare of the person is managed while in custody.

A structured case management process can help to ensure the welfare of the person is cared for.

Additional requirements include:
1. Timing and "stop the clock" for breaks and wait time
1. Involvement of support person, interpreter, etc.
1. Linking to any incident that may occur during the period of custody


classDiagram
  Animal <|-- Duck : LabelText
  class1 --o other
  Animal --o Fish
  Animal : +int age
  Animal : +String gender
  Animal: mate()
  Animal : #method(param)* return
  class Duck{
      %% Class Members
      +String beakColor
      #quack()
  }
  class Fish{
      -abstractMethod()*
      staticMethod()$
  }
  %% Class member generics
  class Square~Shape~{
      List~int~ position
      setPoints(List~int~ points)
      getPoints() List~int~
  }
  Square : -List~string~ messages
  Square : ~setMessages(List~string~ messages)
  Square : +getMessages() List~string~

  %% Multiplicity relations
  Customer "1" --> "*" Ticket
  Student "1" --> "1..*" Course
  Galaxy --> "6" Star : Contains

  %% Annotations
  class Annotate1
  <<interface>> Animal

  class Annotate2{
    <<Service>>
  }