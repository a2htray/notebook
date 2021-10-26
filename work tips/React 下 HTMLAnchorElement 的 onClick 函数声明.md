React 下 HTMLAnchorElement 的 onClick 函数声明
=========================

```js
import * as React from 'react'

export interface ExampleProps {
}

export interface ExampleState {
}

export default class Example extends React.Component<ExampleProps, ExampleState> {
  constructor(props: ExampleProps) {
    super(props);

    this.state = {
    }
  }

  public render() {
    return (
      <div>
        <a href="#" onClick={this.onAnchorClick}>an anchor</a>
      </div>
    )
  }

  onAnchorClick: React.MouseEventHandler<HTMLAnchorElement> = (e: React.MouseEvent<HTMLAnchorElement>) => {
    console.log((e.target as HTMLAnchorElement).innerText)
  }
}
```

```js
onAnchorClick: React.MouseEventHandler<HTMLAnchorElement> = (e: React.MouseEvent<HTMLAnchorElement>) => {
  // ...
}
```