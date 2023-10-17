import { GridOn } from "@mui/icons-material";
import AssignmentIndIcon from "@mui/icons-material/AssignmentInd";
import EditNoteIcon from "@mui/icons-material/EditNote";
import HomeIcon from "@mui/icons-material/Home";
import LoginIcon from "@mui/icons-material/Login";
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
import { CookieSetOptions } from "universal-cookie";

type Props = {
  drawerState: boolean;
  setDrawerState: React.Dispatch<React.SetStateAction<boolean>>;
  removeCookie: (name: "auth", options?: CookieSetOptions | undefined) => void;
};

export default function DrawerComponent({
  drawerState,
  setDrawerState,
  removeCookie,
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
          <Divider />
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
        </List>
      </Box>
    </Drawer>
  );
}
