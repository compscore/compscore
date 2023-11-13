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
import { cookies } from "../models/Cookies";

type Props = {
  cookies: cookies;
  setDrawerState: React.Dispatch<React.SetStateAction<boolean>>;
  mobile: boolean;
};

export default function NavBar({ cookies, setDrawerState, mobile }: Props) {
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
              window.location.href = "/scoreboard";
            }}
            color='inherit'
            sx={{
              textTransform: "none",
            }}
          >
            <Typography variant={mobile ? "body1" : "h6"}>Compscore</Typography>
          </Button>
          <Box sx={{ flexGrow: 1 }}></Box>
          <Box
            sx={{
              display: "flex",
              flexDirection: "row",
              alignItems: "center",
              width: "fit-content",
            }}
          >
            {cookies.auth && (
              <>
                <Typography variant={mobile ? "body1" : "h6"}>
                  {(jwt_decode(cookies.auth) as JWT).Team}
                </Typography>
                {!mobile && (
                  <>
                    <Divider orientation='vertical' flexItem sx={{ m: 1 }} />
                    <Typography variant={mobile ? "body1" : "h6"}>
                      {(jwt_decode(cookies.auth) as JWT).Role}
                    </Typography>
                    <Divider orientation='vertical' flexItem sx={{ m: 1 }} />
                  </>
                )}
              </>
            )}
            <Button href='/'>
              <Avatar src='/compscore.svg' />
            </Button>
          </Box>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
