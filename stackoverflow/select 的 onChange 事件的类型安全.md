select 的 onChange 事件的类型安全
====================================

[原问题](https://stackoverflow.com/questions/33256274/typesafe-select-onchange-event-using-reactjs-and-typescript)

### 问题

我正在尝试在 select 元素上应用 onChange 事件，总是以将其转换成 any 类型，这十分不友好。

有没有可能在类型安全的情况下，获得需要的值。

```typescript jsx
import React = require('react');

interface ITestState {
    selectedValue: string;
}

export class Test extends React.Component<{}, ITestState> {

    constructor() {
        super();
        this.state = { selectedValue: "A" };
    }

    change(event: React.FormEvent) {
        console.log("Test.change");
        console.log(event.target); // in chrome => <select class="form-control" id="searchType" data-reactid=".0.0.0.0.3.1">...</select>

        // Use cast to any works but is not type safe
        var unsafeSearchTypeValue = ((event.target) as any).value;

        console.log(unsafeSearchTypeValue); // in chrome => B

        this.setState({
            selectedValue: unsafeSearchTypeValue
        });
    }

    render() {
        return (
            <div>
                <label htmlFor="searchType">Safe</label>
                <select className="form-control" id="searchType" onChange={ e => this.change(e) } value={ this.state.selectedValue }>
                    <option value="A">A</option>
                    <option value="B">B</option>
                </select>
                <h1>{this.state.selectedValue}</h1>
            </div>
        );
    }
}
```

### 回答

```typescript jsx
import React = require('react');

interface ITestState {
    selectedValue: string;
}

export class Test extends React.Component<{}, ITestState> {

    constructor() {
        super();
        this.state = { selectedValue: "A" };
    }

    change(event: React.FormEvent<HTMLSelectElement>) {
        // No longer need to cast to any - hooray for react!
        var safeSearchTypeValue: string = event.currentTarget.value;

        console.log(safeSearchTypeValue); // in chrome => B

        this.setState({
            selectedValue: safeSearchTypeValue
        });
    }

    render() {
        return (
            <div>
                <label htmlFor="searchType">Safe</label>
                <select className="form-control" id="searchType" onChange={ e => this.change(e) } value={ this.state.selectedValue }>
                    <option value="A">A</option>
                    <option value="B">B</option>
                </select>
                <h1>{this.state.selectedValue}</h1>
            </div>
        );
    }
}
```

### 总结

`event.currentTarget.value`
