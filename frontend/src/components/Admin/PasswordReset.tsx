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
import PasswordInput from "../PasswordInput";

export default function PasswordReset() {
  const [users, setUsers] = useState<[Team]>();
  const [selectedUser, setSelectedUser] = useState<number>();
  const [password, setPassword] = useState<string>();
  const [confirmPassword, setConfirmPassword] = useState<string>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`/api/teams`, {
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

  const changePassword = (password: string) => {
    fetch(`/api/admin/password`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        team: users?.[selectedUser as number].name,
        password: password,
      }),
    })
      .then(async (res) => {
        if (res.status === 200) {
          enqueueSnackbar("Password changed", {
            variant: "success",
          });
        } else {
          enqueueSnackbar("Failed to change password", {
            variant: "error",
          });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Failed to change password", {
          variant: "error",
        });
        console.log(err);
      });
  };

  return (
    <Container maxWidth='xs'>
      <Typography variant='h4' align='center'>
        Password Reset
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
      <PasswordInput
        label='Password'
        margin='normal'
        onChange={(e) => {
          setPassword(e.target.value);
        }}
      />
      <PasswordInput
        label='Confirm Password'
        margin='normal'
        onChange={(e) => {
          setConfirmPassword(e.target.value);
        }}
      />
      {password && confirmPassword && password !== confirmPassword && (
        <Typography color='error' align='center'>
          Passwords do not match
        </Typography>
      )}
      <Box sx={{ m: 1 }} />
      <Button
        variant='contained'
        fullWidth
        onClick={() => {
          if (password && confirmPassword && password === confirmPassword) {
            changePassword(password);
          } else {
            enqueueSnackbar("Password is empty or do not match do not match", {
              variant: "error",
            });
          }
        }}
      >
        <Typography>Change Password</Typography>
      </Button>
    </Container>
  );
}
