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
import { Team } from "../../models/ent";
import { LoginSuccess, LoginFailure } from "../../models/Login";
import { CookieSetOptions } from "universal-cookie";

type Props = {
  setCookie: (
    name: "auth",
    value: any,
    options?: CookieSetOptions | undefined
  ) => void;
};

export default function AuthenticateAs({ setCookie }: Props) {
  const [users, setUsers] = useState<[Team]>();
  const [selectedUser, setSelectedUser] = useState<number>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`http://localhost:8080/api/teams`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            let response = (await res.json()) as [Team];

            setUsers(response);
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
    fetch(`http://localhost:8080/api/admin/login`, {
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
        console.log(res);
        if (res.status === 200) {
          let response = (await res.json()) as LoginSuccess;

          setCookie("auth", response.token, {
            path: response.path,
            domain: response.domain,
            secure: response.secure,
            httpOnly: response.httpOnly,
            expires: new Date(response.expiration * 1000),
          });

          window.location.href = "/";
        } else {
          let response = (await res.json()) as LoginFailure;

          enqueueSnackbar(response.error, { variant: "error" });
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
        <Typography>Authenticate As</Typography>
      </Button>
    </Container>
  );
}
