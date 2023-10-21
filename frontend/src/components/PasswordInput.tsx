import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import IconButton from "@mui/material/IconButton";
import InputAdornment from "@mui/material/InputAdornment";
import TextField from "@mui/material/TextField";
import { useState } from "react";

type props = {
  value?: string;
  onBlur?: React.ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;
  variant?: "standard" | "filled" | "outlined" | undefined;
};

function PasswordInput({ value, variant, onBlur }: props) {
  const [password, setPassword] = useState<string>(
    value === undefined ? "" : value
  );
  const [prevPassword, setPrevPassword] = useState<string>(
    value === undefined ? "" : value
  );
  const [showPassword, setShowPassword] = useState(false);

  const handleTogglePassword = () => {
    setShowPassword((prevShowPassword) => !prevShowPassword);
  };

  const handleBlur = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    if (onBlur && password !== prevPassword) {
      onBlur(e);
      setPrevPassword(password);
    }
  };

  return (
    <TextField
      type={showPassword ? "text" : "password"}
      variant={variant}
      value={password}
      onChange={(e) => {
        setPassword(e.target.value);
      }}
      onBlur={handleBlur}
      fullWidth
      InputProps={{
        endAdornment: (
          <InputAdornment position='end'>
            <IconButton onClick={handleTogglePassword}>
              {showPassword ? <VisibilityOffIcon /> : <VisibilityIcon />}
            </IconButton>
          </InputAdornment>
        ),
      }}
    />
  );
}

export default PasswordInput;
