import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { enqueueSnackbar } from "notistack";
import { CheckScoreboard } from "../models/CheckScoreboard";
import {
  Box,
  Typography,
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
} from "@mui/material";

export default function CheckScoreBoard() {
  const { check: check } = useParams() as { check: string };
  const [data, setData] = useState<CheckScoreboard>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`http://localhost:8080/api/scoreboard/check/${check}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          let response = (await res.json()) as CheckScoreboard;
          if (res.status === 200) {
            setData(response);
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

    console.log(data);
    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedRound, setHighlightedRound] = useState<number | null>(null);
  const [highlightedTeam, setHighlightedTeam] = useState<string | null>(null);

  const getBackgroundColor = (status: number, round: number, check: string) => {
    if (highlightedTeam === null && highlightedRound === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedTeam === check || highlightedRound === round) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (status === 0) {
      return "#f26359";
    } else if (status === 1) {
      return "#75ad77";
    }
    return "#cecdc6";
  };

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
        {check}
      </Typography>
      <Typography component='h1' variant='h5'>
        Round {data?.round}
      </Typography>
      <Box m={2}></Box>
      <TableContainer component={Paper} sx={{ width: "80%" }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell size='small'>Round</TableCell>
              {data?.teams[0].status.map((_, index) => (
                <TableCell
                  size='small'
                  key={"round-" + (data?.round - index)}
                  sx={{
                    backgroundColor:
                      highlightedRound === index ? "#343434" : "transparent",
                  }}
                  onMouseEnter={() => {
                    setHighlightedRound(index);
                  }}
                  onMouseLeave={() => {
                    setHighlightedRound(null);
                  }}
                >
                  {data?.round - index}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.teams.map((team, index) => (
              <TableRow key={index}>
                <TableCell
                  size='small'
                  sx={{
                    backgroundColor:
                      highlightedTeam === team.name ? "#343434" : "transparent",
                  }}
                  onMouseEnter={() => {
                    setHighlightedTeam(team.name);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                  onClick={() => {
                    window.location.href = "/scoreboard/team/" + (index + 1);
                  }}
                >
                  {team.name}
                </TableCell>
                {team.status.map((status, index) => (
                  <TableCell
                    key={index}
                    size='small'
                    sx={{
                      backgroundColor: getBackgroundColor(
                        status,
                        index,
                        team.name
                      ),
                    }}
                    onMouseEnter={() => {
                      setHighlightedRound(index);
                      setHighlightedTeam(team.name);
                    }}
                    onMouseLeave={() => {
                      setHighlightedRound(null);
                      setHighlightedTeam(null);
                    }}
                  ></TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
