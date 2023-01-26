import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import { Link } from "react-router-dom";
import Drawer from "@mui/material/Drawer";
import { useState } from 'react';
import Divider from "@mui/material/Divider";

function Navbar() {
    const [drawerOpen, setDrawerOpen] = useState(false)

    return (
        <Box sx={{ flexGrow: 1}}>
            <AppBar position="static" >
                <Toolbar>
                    <IconButton
                        size="large"
                        edge="start"
                        color="inherit"
                        aria-label="menu"
                        sx={{ mr: 2 }}
                        onClick={() => setDrawerOpen(true)}
                    >
                        <MenuIcon />
                    </IconButton>
                    <Drawer
                        anchor='left'
                        open={drawerOpen}
                        onClose={() => setDrawerOpen(false)}
                    >
                        <Box p={2} width='200px' textAlign='center'>
                            <Typography variant="h6" component="div">
                                <Divider />
                                <Link to="/" onClick={() => setDrawerOpen(false)}>
                                    Homepage
                                </Link>
                                <Divider />
                            </Typography>
                            <Typography variant="h6" component="div">
                                <Link to="/treatment_grid" onClick={() => setDrawerOpen(false)}>
                                    Treatment
                                </Link>
                                <Divider />
                            </Typography>
                            <Typography variant="h6" component="div">
                                <Link to="/treatment_plan_grid" onClick={() => setDrawerOpen(false)}>
                                    Treatment Plan
                                </Link>
                                <Divider />
                            </Typography>
                            
                        </Box>
                    </Drawer>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                        <Link className="app-header" to="/">
                            ระบบทันตกรรม
                        </Link>
                    </Typography>
                </Toolbar>
            </AppBar>
        </Box>
    );
}

export default Navbar;
