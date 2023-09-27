import React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import { enqueueSnackbar } from "notistack";
import { LoginSuccess, LoginFailure } from "../models/Login";
import { CookieSetOptions } from "universal-cookie";

type Props = {
  setCookie: (
    name: "auth",
    value: any,
    options?: CookieSetOptions | undefined
  ) => void;
};

export default function Login({ setCookie }: Props) {
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = JSON.stringify({
      username: event.currentTarget.username.value,
      password: event.currentTarget.password.value,
    });

    fetch("http://localhost:8080/api/login", {
      method: "POST",
      body: data,
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        let response = await res.json();
        if (res.status === 200) {
          enqueueSnackbar("Logged in", { variant: "success" });

          response = response as LoginSuccess;

          setCookie("auth", response.token, {
            path: response.path,
            domain: response.domain,
            secure: response.secure,
            httpOnly: response.httpOnly,
            expires: response.expires,
          });

          window.location.href = "/";
        } else {
          response = response as LoginFailure;

          enqueueSnackbar(response.error, { variant: "error" });
        }
        return res.json();
      })
      .catch((err) => {
        enqueueSnackbar("Encountered an error" + err, { variant: "error" });
        console.log(err);
      });
  };

  return (
    <Box
      sx={{
        marginTop: 8,
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography component='h1' variant='h5'>
        Sign in
      </Typography>
      <Box component='form' onSubmit={handleSubmit}>
        <TextField
          margin='normal'
          required
          fullWidth
          id='username'
          label='Team'
          name='username'
          autoFocus
        />
        <TextField
          margin='normal'
          required
          fullWidth
          name='password'
          label='Password'
          type='password'
          id='password'
        />
        <Button
          type='submit'
          fullWidth
          variant='contained'
          sx={{ mt: 3, mb: 2 }}
        >
          Sign In
        </Button>
      </Box>
    </Box>
  );
}
