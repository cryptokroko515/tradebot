'use strict';

import React from 'react';

import AppBar from 'material-ui/AppBar';
import Drawer from './Drawer';
import { withRouter } from 'react-router-dom';


class Header extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			drawer: false,
		};
		this.handleDrawerToggle = this.handleDrawerToggle.bind(this);
		this.handleDrawerChange = this.handleDrawerChange.bind(this);
		this.handleTitleTap = this.handleTitleTap.bind(this);
	}

	handleDrawerToggle() {
		this.setState({ drawer: ! this.state.drawer });
	}

	handleDrawerChange(status) {
		this.setState({ drawer: status });
	}

	handleTitleTap() {
		this.props.history.push('/');
	}


	render() {

		return (
			<div className="component--appbar">
				<AppBar
					title={ this.props.title || 'Tradebot' }
					onLeftIconButtonTouchTap={ this.handleDrawerToggle }
					onTitleTouchTap={ this.handleTitleTap }
				/>
				<Drawer open={ this.state.drawer } change={ this.handleDrawerChange } />
			</div>
		)
	}

}

export default withRouter( Header );
