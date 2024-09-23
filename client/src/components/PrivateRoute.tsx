import React from 'react';
import { Navigate } from 'react-router-dom';
import { useUser } from './UserContext';

interface PrivateRouteProps {
    children: JSX.Element;
}

export const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
    const { username } = useUser();

    if (!username) {
        return <Navigate to="/" />;
    }

    // If user is authenticated, render the page
    return children;
};

