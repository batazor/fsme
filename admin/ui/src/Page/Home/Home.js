import React from 'react'
import Typography from '@material-ui/core/Typography'

export default () => {
  return (
    <div>
      <Typography variant="h2" color="inherit">
        FSME-UI
      </Typography>

      <Typography variant="h6" color="inherit">Link: </Typography>
      <Typography variant="body1" color="inherit">
        + GitHub repository: <a href="https://github.com/batazor/fsme">https://github.com/batazor/fsme</a>
      </Typography>
      <Typography variant="body1" color="inherit">
        + Version: 0.1.0
      </Typography>
    </div>
  );
}
