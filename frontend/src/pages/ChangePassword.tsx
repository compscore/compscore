import { Box, Container, Typography } from "@mui/material";
import PasswordInput from "../components/PasswordInput";

export default function ChangePassword() {
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
        <Box sx={{ m: 2 }} />
        <PasswordInput />
        <Box sx={{ m: 2 }} />
        <PasswordInput />
        <Box sx={{ m: 2 }} />
        <PasswordInput />
      </Box>
    </Container>
  );
}
