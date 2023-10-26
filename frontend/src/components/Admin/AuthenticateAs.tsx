import {
  Box,
  Button,
  Container,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { api_url, domain, path } from "../../config";
import { setCookie, cookies } from "../../models/Cookies";
import { LoginFailure, LoginSuccess } from "../../models/Login";
import { Team } from "../../models/ent";

type Props = {
  setCookie: setCookie;
  cookies: cookies;
};

export default function AuthenticateAs({ setCookie, cookies }: Props) {
  const [users, setUsers] = useState<[Team]>();
  const [selectedUser, setSelectedUser] = useState<number>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`${api_url}/api/teams`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            setUsers((await res.json()) as [Team]);
          } else {
            enqueueSnackbar("Encountered an error", { variant: "error" });
          }
        })
        .catch((err) => {
          enqueueSnackbar("Encountered an error: " + err, { variant: "error" });
          console.log(err);
        });
    };

    fetchData();
  }, []);

  const authenticateAs = () => {
    fetch(`${api_url}/api/admin/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        team: users?.[selectedUser as number].name,
      }),
    })
      .then(async (res) => {
        if (res.status === 200) {
          setCookie("admin", cookies.auth, { path: path, domain: domain });
          setCookie("auth", ((await res.json()) as LoginSuccess).token, {
            path: path,
            domain: domain,
          });

          window.location.href = "/";
        } else {
          enqueueSnackbar(((await res.json()) as LoginFailure).error, {
            variant: "error",
          });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Encountered an error" + err, { variant: "error" });
        console.log(err);
      });
  };

  const loginAs = () => {
    fetch(`${api_url}/api/admin/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        team: users?.[selectedUser as number].name,
      }),
    })
      .then(async (res) => {
        if (res.status === 200) {
          const response = (await res.json()) as LoginSuccess;

          setCookie("auth", response.token, {
            path: path,
            domain: domain,
            secure: response.secure,
            httpOnly: response.httpOnly,
            expires: new Date(response.expiration * 1000),
          });

          window.location.href = "/";
        } else {
          enqueueSnackbar(((await res.json()) as LoginFailure).error, {
            variant: "error",
          });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Encountered an error" + err, { variant: "error" });
        console.log(err);
      });
  };

  return (
    <Container maxWidth='xs'>
      <Typography variant='h4' align='center'>
        Authenticate As
      </Typography>
      <Box sx={{ m: 2 }} />
      <FormControl fullWidth margin='normal'>
        <InputLabel id='user-select-label'>User</InputLabel>
        <Select
          labelId='user-select-label'
          value={selectedUser}
          label='Age'
          onChange={(e) => {
            setSelectedUser(e.target.value as number);
          }}
        >
          <MenuItem value={undefined}>
            <em>None</em>
          </MenuItem>
          {users?.map((user, index) => (
            <MenuItem key={user.name} value={index}>
              {user.name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
      <Button
        variant='contained'
        fullWidth
        onClick={() => {
          authenticateAs();
        }}
      >
        <Typography>Temporary Authentication</Typography>
      </Button>
      <Box sx={{ m: 2 }} />
      <Button
        variant='contained'
        fullWidth
        onClick={() => {
          loginAs();
        }}
      >
        <Typography>Full Authentication</Typography>
      </Button>
    </Container>
  );
}
