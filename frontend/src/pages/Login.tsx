import {
  Box,
  Button,
  Container,
  Link,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";
import React from "react";
import PasswordInput from "../components/PasswordInput";
import { setCookie } from "../models/Cookies";
import { LoginFailure, LoginSuccess } from "../models/Login";
import { api_url, domain, path } from "../config";

type props = {
  setCookie: setCookie;
};

export default function Login({ setCookie }: props) {
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = JSON.stringify({
      username: event.currentTarget.username.value,
      password: event.currentTarget.password.value,
    });

    fetch(`${api_url}/api/login`, {
      method: "POST",
      body: data,
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        if (res.status === 200) {
          const response = (await res.json()) as LoginSuccess;
          enqueueSnackbar("Logged in", { variant: "success" });

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
    <Container component='main' maxWidth='xs'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3'>
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
          <PasswordInput
            margin='normal'
            required
            name='password'
            label='Password'
            id='password'
          />
          <Tooltip title='Reach out to Black Team' arrow>
            <Link
              sx={{
                textDecoration: "none",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
              }}
            >
              <Typography variant='body2' color='text.secondary'>
                Forgot password?
              </Typography>
            </Link>
          </Tooltip>
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
    </Container>
  );
}
