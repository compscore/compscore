import { useState, useEffect } from "react";
import { Scoreboard } from "../models/Scoreboard";
import { enqueueSnackbar } from "notistack";
import { Box, Typography } from "@mui/material";

export default function ScoreBoard() {
  const [data, setData] = useState<Scoreboard>();

  useEffect(() => {
    const fetchData = async () => {
      fetch("http://localhost:8080/api/scoreboard", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }).then(async (res) => {
        let response = (await res.json()) as Scoreboard;
        if (res.status === 200) {
          console.log(response);
          setData(response);
        } else {
          enqueueSnackbar("Encountered an error", { variant: "error" });
        }
      });
    };

    fetchData();

    const pollingInterval = setInterval(fetchData, 5000);

    return () => clearInterval(pollingInterval);
  }, []);

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography
        component='h1'
        variant='h3'
        fontWeight={700}
        sx={{
          marginTop: 5,
        }}
      >
        Scoreboard
      </Typography>
      <Typography component='h1' variant='h5'>
        Round {data?.round}
      </Typography>
      {data?.checks.map((check) => (
        <>
          <Typography component='h1' variant='h5'>
            {check.name}
          </Typography>
          <Box
            sx={{
              display: "flex",
              flexDirection: "row",
              alignItems: "center",
            }}
          >
            {check.teams.map((team) => (
              <Typography component='h1' variant='h5'>
                {team}
              </Typography>
            ))}
          </Box>
        </>
      ))}
    </Box>
  );
}
