// @flow
import * as React from 'react'
import {Box2, Text} from '../../../common-adapters'
import {styleSheetCreate} from '../../../styles'

type Props = {
  skeleton: null,
}

const Available = ({skeleton}: Props) => (
  <Box2 direction="vertical">
    <Text type="Body" style={styles.text}>
      Available {skeleton}
    </Text>
  </Box2>
)

const styles = styleSheetCreate({
  text: {
    textAlign: 'center',
  },
})

export default Available
