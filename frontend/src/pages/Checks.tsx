import {
  Box,
  Container,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import PasswordInput from "../components/PasswordInput";
import { useEffect, useState } from "react";
import { credentials } from "../models/Credentials";
import { enqueueSnackbar } from "notistack";

type Props = {
  cookies: {
    auth?: any;
  };
};

export default function Checks({ cookies }: Props) {
  if (cookies.auth == undefined) {
    window.location.href = "/login";
  }

  const [credentials, setCredentials] = useState<credentials | undefined>(
    undefined
  );
  const fetchChecks = () => {
    fetch("http://localhost:8080/api/checks", {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        let response = await res.json();
        if (res.status === 200) {
          setCredentials(response);
        } else {
          console.log(response);
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    fetchChecks();
  }, []);

  const updatePassword = (password: string, check: string) => {
    fetch(`http://localhost:8080/api/credential/${check}`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        password: password,
      }),
    }).then(async (res) => {
      if (res.status === 200) {
        enqueueSnackbar("Updated password for " + check, {
          variant: "success",
        });
      } else {
        enqueueSnackbar("Failed to update password for " + check, {
          variant: "error",
        });
      }
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
        <Typography component='h1' variant='h3' fontWeight={700}>
          Check Editor
        </Typography>
        <Box sx={{ m: 1 }} />
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>
                  <Typography>Check Name</Typography>
                </TableCell>
                <TableCell>
                  <Typography>Password</Typography>
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {credentials?.map((credential) => (
                <TableRow>
                  <TableCell>
                    <Typography variant='body1' component='h1'>
                      {credential.check}
                    </Typography>
                  </TableCell>

                  <TableCell>
                    <PasswordInput
                      value={credential.password}
                      onBlur={(e) => {
                        updatePassword(e.target.value, credential.check);
                      }}
                    />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Box>
    </Container>
  );
}
