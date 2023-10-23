import { Box, Button, Container, Typography } from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useState } from "react";
import PasswordInput from "../components/PasswordInput";
import { api_url, domain, path } from "../config";
import { removeCookie } from "../models/Cookies";

type props = {
  removeCookie: removeCookie;
};

export default function ChangePassword({ removeCookie }: props) {
  const [oldPassword, setOldPassword] = useState<string>("");
  const [newPassword, setNewPassword] = useState<string>("");
  const [confirmNewPassword, setConfirmNewPassword] = useState<string>("");

  const changePassword = (oldPassword: string, newPassword: string) => {
    fetch(`${api_url}/api/password`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        oldPassword: oldPassword,
        newPassword: newPassword,
      }),
    })
      .then(async (res) => {
        if (res.status === 200) {
          enqueueSnackbar("Password changed", {
            variant: "success",
          });
          removeCookie("auth", { path: path, domain: domain });
          window.location.href = "/login";
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
    <Container component='main' maxWidth='xs'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography variant='h3'>Change Password</Typography>
        <PasswordInput
          label='Old Password'
          margin='normal'
          onChange={(e) => {
            setOldPassword(e.target.value);
          }}
        />
        <PasswordInput
          label='New Password'
          margin='normal'
          onChange={(e) => {
            setNewPassword(e.target.value);
          }}
        />
        <PasswordInput
          label='Confirm New Password'
          margin='normal'
          onChange={(e) => {
            setConfirmNewPassword(e.target.value);
          }}
        />
        {newPassword &&
          confirmNewPassword &&
          newPassword !== confirmNewPassword && (
            <Typography color='error'>Passwords do not match</Typography>
          )}
        <Box sx={{ m: 1 }} />
        <Button
          variant='contained'
          fullWidth
          onClick={() => {
            if (
              newPassword &&
              confirmNewPassword &&
              newPassword === confirmNewPassword
            ) {
              changePassword(oldPassword, newPassword);
            } else {
              enqueueSnackbar("Passwords do not match", { variant: "error" });
            }
          }}
        >
          <Typography>Change Password</Typography>
        </Button>
      </Box>
    </Container>
  );
}
