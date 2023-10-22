import { GridOn } from "@mui/icons-material";
import AdminPanelSettingsIcon from "@mui/icons-material/AdminPanelSettings";
import AssignmentIndIcon from "@mui/icons-material/AssignmentInd";
import EditNoteIcon from "@mui/icons-material/EditNote";
import HomeIcon from "@mui/icons-material/Home";
import LoginIcon from "@mui/icons-material/Login";
import PasswordIcon from "@mui/icons-material/Password";
import SupervisorAccountIcon from "@mui/icons-material/SupervisorAccount";
import {
  Box,
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import jwt_decode from "jwt-decode";
import { cookies, removeCookie, setCookie } from "../models/Cookies";
import { JWT } from "../models/JWT";

type Props = {
  drawerState: boolean;
  setDrawerState: React.Dispatch<React.SetStateAction<boolean>>;
  setCookie: setCookie;
  removeCookie: removeCookie;
  cookies: cookies;
};

export default function DrawerComponent({
  drawerState,
  setDrawerState,
  setCookie,
  removeCookie,
  cookies,
}: Props) {
  const toggleDrawer =
    (open: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
      if (
        event.type === "keydown" &&
        ((event as React.KeyboardEvent).key === "Tab" ||
          (event as React.KeyboardEvent).key === "Shift")
      ) {
        return;
      }

      setDrawerState(open);
    };

  return (
    <Drawer anchor={"left"} open={drawerState} onClose={toggleDrawer(false)}>
      <Box
        sx={{ width: 250 }}
        role='presentation'
        onClick={toggleDrawer(false)}
        onKeyDown={toggleDrawer(false)}
      >
        <List>
          <ListItem
            disablePadding
            onClick={() => {
              window.location.href = "/";
            }}
          >
            <ListItemButton>
              <ListItemIcon>{<HomeIcon />}</ListItemIcon>
              <ListItemText primary='Home' />
            </ListItemButton>
          </ListItem>
          <ListItem
            disablePadding
            onClick={() => {
              window.location.href = "/scoreboard";
            }}
          >
            <ListItemButton>
              <ListItemIcon>{<GridOn />}</ListItemIcon>
              <ListItemText primary='Scoreboard' />
            </ListItemButton>
          </ListItem>
          {cookies.auth ? (
            (jwt_decode(cookies.auth) as JWT).Role === "admin" ? (
              <ListItem
                disablePadding
                onClick={() => {
                  window.location.href = "/admin";
                }}
              >
                <ListItemButton>
                  <ListItemIcon>{<AdminPanelSettingsIcon />}</ListItemIcon>
                  <ListItemText primary='Admin Panel' />
                </ListItemButton>
              </ListItem>
            ) : (
              <ListItem
                disablePadding
                onClick={() => {
                  window.location.href = "/checks";
                }}
              >
                <ListItemButton>
                  <ListItemIcon>{<EditNoteIcon />}</ListItemIcon>
                  <ListItemText primary='Edits Checks' />
                </ListItemButton>
              </ListItem>
            )
          ) : (
            <></>
          )}
          <Divider />
          {cookies.auth ? (
            <>
              {cookies.admin ? (
                <ListItem
                  disablePadding
                  onClick={() => {
                    setCookie("auth", cookies.admin);
                    removeCookie("admin");
                    window.location.href = "/";
                  }}
                >
                  <ListItemButton>
                    <ListItemIcon>{<SupervisorAccountIcon />}</ListItemIcon>
                    <ListItemText primary='Return to Admin' />
                  </ListItemButton>
                </ListItem>
              ) : (
                <ListItem
                  disablePadding
                  onClick={() => {
                    removeCookie("auth");
                    window.location.href = "/";
                  }}
                >
                  <ListItemButton>
                    <ListItemIcon>{<AssignmentIndIcon />}</ListItemIcon>
                    <ListItemText primary='Logout' />
                  </ListItemButton>
                </ListItem>
              )}
              <ListItem
                disablePadding
                onClick={() => {
                  window.location.href = "/password";
                }}
              >
                <ListItemButton>
                  <ListItemIcon>{<PasswordIcon />}</ListItemIcon>
                  <ListItemText primary='Change Password' />
                </ListItemButton>
              </ListItem>
            </>
          ) : (
            <ListItem
              disablePadding
              onClick={() => {
                window.location.href = "/login";
              }}
            >
              <ListItemButton>
                <ListItemIcon>{<LoginIcon />}</ListItemIcon>
                <ListItemText primary='Login' />
              </ListItemButton>
            </ListItem>
          )}
        </List>
      </Box>
    </Drawer>
  );
}
