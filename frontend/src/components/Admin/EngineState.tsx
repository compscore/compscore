import {
  Button,
  ButtonGroup,
  Box,
  Container,
  Typography,
  Tooltip,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { EngineMessage, EngineStatus } from "../../models/Engine";

export default function EngineState() {
  const [engineState, setEngineState] = useState<EngineStatus>({
    status: "unknown",
    message: "Engine state unknown",
  });

  useEffect(() => {
    const fetchData = async () => {
      fetch(`/api/engine`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            let response = (await res.json()) as EngineStatus;

            setEngineState(response);
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

    const pollingInterval = setInterval(fetchData, 5000);

    return () => clearInterval(pollingInterval);
  }, []);

  const getChipColor = (status: EngineStatus) => {
    switch (status.status) {
      case "running":
        return "success";
      case "paused":
        return "error";
      case "unknown":
        return "warning";
      default:
        return "warning";
    }
  };

  const startEngine = () => {
    fetch(`/api/engine/start`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        if (res.status === 200) {
          let response = (await res.json()) as EngineMessage;
          enqueueSnackbar(`Engine started: ${response.message}`, {
            variant: "success",
          });
        } else {
          let response = (await res.json()) as { error: string };

          enqueueSnackbar(`Encountered an error: ${response.error}`, {
            variant: "error",
          });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Encountered an error: " + err, { variant: "error" });
        console.log(err);
      });
  };

  const stopEngine = () => {
    fetch(`/api/engine/stop`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(async (res) => {
        if (res.status === 200) {
          let response = (await res.json()) as EngineMessage;
          enqueueSnackbar(`Engine started: ${response.message}`, {
            variant: "success",
          });
        } else {
          let response = (await res.json()) as { error: string };

          enqueueSnackbar(`Encountered an error: ${response.error}`, {
            variant: "error",
          });
        }
      })
      .catch((err) => {
        enqueueSnackbar("Encountered an error: " + err, { variant: "error" });
        console.log(err);
      });
  };

  return (
    <Container maxWidth='xs'>
      <Typography variant='h4' align='center'>
        Engine State
      </Typography>
      <Box
        sx={{ m: 2 }}
        display='flex'
        alignItems='center'
        flexDirection='column'
      >
        <Tooltip title={engineState.message}>
          <Button variant='contained' color={getChipColor(engineState)}>
            <Typography variant='h5'>{engineState.status}</Typography>
          </Button>
        </Tooltip>

        <Box sx={{ m: 2 }} />
        <ButtonGroup variant='contained' fullWidth>
          <Button onClick={startEngine}>
            <Typography variant='h6'>Start</Typography>
          </Button>
          <Button onClick={stopEngine}>
            <Typography variant='h6'>Pause</Typography>
          </Button>
        </ButtonGroup>
      </Box>
    </Container>
  );
}
