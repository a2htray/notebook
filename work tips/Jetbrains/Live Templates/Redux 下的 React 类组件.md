key: react_cls_component_redux

```typescript jsx
import React from 'react'
import {connect} from 'react-redux'

const mapStateToProps = (state: any) => ({
  
})

const dispatchProps = {
  
}

type Props = ReturnType<typeof mapStateToProps> & typeof dispatchProps & {
  
}

type State = {
  
}

class $ClassName$ extends React.Component<Props, State>{
  render() {
    return <>
      
    </>;
  }
}

export default connect(mapStateToProps, dispatchProps)($ClassName$)
```