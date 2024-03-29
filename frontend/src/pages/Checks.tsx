import {
  Box,
  Container,
  Paper,
  Switch,
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
import { Credential } from "../models/ent";
import { enqueueSnackbar } from "notistack";
import jwt_decode from "jwt-decode";
import { JWT } from "../models/JWT";
import { cookies } from "../models/Cookies";
import { api_url, fetchWithTimeout } from "../config";

type Props = {
  cookies: cookies;
};

export default function Checks({ cookies }: Props) {
  const [showPasswordlessChecks, setShowPasswordlessChecks] =
    useState<boolean>(false);
  const [credentials, setCredentials] = useState<[Credential] | undefined>(
    undefined
  );

  const fetchChecks = () => {
    fetchWithTimeout(`${api_url}/api/credentials`, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        const response = (await res.json()) as [Credential];
        if (res.status === 200) {
          setCredentials(response);
        } else {
          enqueueSnackbar("Failed to fetch checks", { variant: "error" });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Failed to fetch checks: " + err, {
          variant: "error",
        });
        console.log(err);
      });
  };

  useEffect(() => {
    fetchChecks();
  }, []);

  if (cookies.auth == undefined) {
    window.location.href = "/login";
  }

  if (cookies.auth && (jwt_decode(cookies.auth) as JWT).Role === "admin") {
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
          <Typography variant='h6' component='h1'>
            You do not have any checks to edit.
          </Typography>
        </Box>
      </Container>
    );
  }

  const updatePassword = (password: string, check: string) => {
    fetchWithTimeout(`${api_url}/api/credential/${check}`, {
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
        <Box
          sx={{ m: 1 }}
          display='flex'
          flexDirection='row'
          alignItems='center'
        >
          <Switch
            onChange={(e) => {
              setShowPasswordlessChecks(e.target.checked);
            }}
          />
          <Typography>Show checks without passwords</Typography>
        </Box>
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
              {credentials
                ?.filter(
                  (credential) =>
                    credential.password !== "" || showPasswordlessChecks
                )
                .map((credential) => (
                  <TableRow key={credential.edges.check?.name}>
                    <TableCell>
                      <Typography variant='body1' component='h1'>
                        {credential.edges.check?.name}
                      </Typography>
                    </TableCell>

                    <TableCell>
                      <PasswordInput
                        value={credential.password}
                        variant='standard'
                        onBlur={(e) => {
                          if (credential.edges.check) {
                            updatePassword(
                              e.target.value,
                              credential.edges.check?.name
                            );
                          }
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
