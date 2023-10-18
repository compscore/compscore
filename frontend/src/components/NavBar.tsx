import MenuIcon from "@mui/icons-material/Menu";
import {
  AppBar,
  Avatar,
  Box,
  Button,
  Divider,
  Toolbar,
  Typography,
} from "@mui/material";
import jwt_decode from "jwt-decode";
import { JWT } from "../models/JWT";

type Props = {
  cookies: {
    auth?: any;
  };
  setDrawerState: React.Dispatch<React.SetStateAction<boolean>>;
};

export default function NavBar({ cookies, setDrawerState }: Props) {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position='static'>
        <Toolbar>
          <Button
            color='inherit'
            onClick={() => {
              setDrawerState(true);
            }}
          >
            <MenuIcon />
          </Button>
          <Box sx={{ m: 1 }} />
          <Button
            onClick={() => {
              window.location.href = "/";
            }}
            color='inherit'
            sx={{
              textTransform: "none",
            }}
          >
            <Typography variant='h6'>Compscore</Typography>
          </Button>
          <Box sx={{ flexGrow: 1 }}></Box>
          {cookies.auth && (
            <Box
              sx={{
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                width: "fit-content",
              }}
            >
              <Typography variant='h6'>
                {(jwt_decode(cookies.auth) as JWT).Team}
              </Typography>
              <Box sx={{ m: 1 }} />
              <Divider orientation='vertical' flexItem />
              <Box sx={{ m: 1 }} />
              <Typography variant='h6'>
                {(jwt_decode(cookies.auth) as JWT).Role}
              </Typography>
              <Box sx={{ m: 1 }} />
              <Divider orientation='vertical' flexItem />
            </Box>
          )}
          <Box sx={{ m: 1 }} />
          <Button href='/'>
            <Avatar src='/compscore.svg' />
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
