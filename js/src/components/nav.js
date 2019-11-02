import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import MenuIcon from '@material-ui/icons/Menu';
import {makeStyles } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import MenuItem from '@material-ui/core/MenuItem';
import RoutesConfig from './routes_config';
import Toolbar from '@material-ui/core/Toolbar';
import Menu from '@material-ui/core/Menu';

const useStyles = makeStyles(theme => ({
  menuButton: {
    marginRight: theme.spacing(2),
  },
}));

function MenuItems (props){
	const classes = useStyles();
	const [anchorEl, setAnchorEl] = React.useState(null);

	const handleClick = event => {
    	setAnchorEl(event.currentTarget);
  	};

  	const handleClose = () => {
 		setAnchorEl(null);
  	};
  	return (
  		<div>
			<IconButton
				onClick={handleClick}
				edge="start"
				className={classes.menuButton}
				color="inherit"
				aria-label="open drawer">
				<MenuIcon />
			</IconButton>
			<Menu
				id="simple-menu"
				anchorEl={anchorEl}
				keepMounted
				open={Boolean(anchorEl)}
				onClose={handleClose}>
				{RoutesConfig.map((c) => {
					return (
						<a href={c.route}>{c.label}</a>
					)	
				})}	        	
			</Menu>
  		</div>
	)
}
export default function(){
	return (
		<AppBar position="static">
			<Toolbar>
				<MenuItems />
			</Toolbar>
		</AppBar>
	)
}