import { useState, useEffect } from "react";
import { Scoreboard } from "../models/Scoreboard";
import { enqueueSnackbar } from "notistack";
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

export default function ScoreBoard() {
  const [data, setData] = useState<Scoreboard>();

  useEffect(() => {
    const fetchData = async () => {
      fetch("http://localhost:8080/api/scoreboard", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          let response = (await res.json()) as Scoreboard;
          if (res.status === 200) {
            console.log(response);
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

    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedTeam, setHighlightedTeam] = useState<number | null>(null);
  const [highlightedCheck, setHighlightedCheck] = useState<string | null>(null);

  const getBackgroundColor = (status: number, team: number, check: string) => {
    if (highlightedCheck === null && highlightedTeam === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedCheck === check || highlightedTeam === team) {
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
        Scoreboard
      </Typography>
      <Typography component='h1' variant='h5'>
        Round {data?.round}
      </Typography>
      <Box m={2}></Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell size='small'>
                <Typography variant='subtitle2'>Team</Typography>
              </TableCell>
              {data?.checks[0].teams.map((_, index) => (
                <TableCell
                  key={index + 1}
                  align='center'
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedTeam(index + 1);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                >
                  <Typography variant='subtitle2'>{index + 1}</Typography>
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.checks.map((check) => (
              <TableRow>
                <TableCell
                  key={check.name}
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedCheck(check.name);
                  }}
                  onMouseLeave={() => {
                    setHighlightedCheck(null);
                  }}
                >
                  {check.name}
                </TableCell>
                {check.teams.map((status, team) => (
                  <TableCell
                    key={team + "-" + check.name}
                    size='small'
                    sx={{
                      backgroundColor: getBackgroundColor(
                        status,
                        team + 1,
                        check.name
                      ),
                    }}
                    onMouseEnter={() => {
                      setHighlightedTeam(team + 1);
                      setHighlightedCheck(check.name);
                    }}
                    onMouseLeave={() => {
                      setHighlightedTeam(null);
                      setHighlightedCheck(null);
                    }}
                  ></TableCell>
                ))}
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Box m={2}></Box>
    </Box>
  );
}