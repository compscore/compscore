import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import { Avatar } from "@mui/material";
import Link from "@mui/material/Link";

export default function ButtonAppBar() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static'>
        <Toolbar>
          <Button href='/'>
            <Avatar src='/compscore.svg' />
          </Button>
          <Box sx={{ m: 1 }} />
          <Link
            href='/'
            color='inherit'
            underline='none'
            variant='h6'
            sx={{ flexGrow: 1 }}
          >
            Compscore
          </Link>
          <Button color='inherit' href='/login'>
            Login
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
