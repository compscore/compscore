import { Box } from "@mui/material";
import { useState, useEffect } from "react";
import { Scoreboard } from "../models/Scoreboard";
import { enqueueSnackbar } from "notistack";

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
      <h1>Scoreboard</h1>
      <h2>Round {data?.round}</h2>
    </Box>
  );
}
