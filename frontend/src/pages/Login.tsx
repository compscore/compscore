import React from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";

export default function Login() {
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = JSON.stringify({
      username: event.currentTarget.username.value,
      password: event.currentTarget.password.value,
    });

    console.log(data);

    let response = await fetch("http://localhost:8080/api/login", {
      method: "POST",
      body: data,
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((data) => console.log(data))
      .catch((err) => console.log(err));
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
