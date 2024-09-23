import { createContext, useState, ReactNode, useContext } from 'react';


interface UserContextType {
    username: string | null;
    setUser: (user: string | null) => void;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider = ({ children }: { children: ReactNode }) => {
    const [username, setUser] = useState<string | null>(null);
    return (
        <UserContext.Provider value={{ username, setUser }}>
            {children}
        </UserContext.Provider>
    );
};

export const useUser = () => {
    const context = useContext(UserContext);
    if (context === undefined) {
        // TODO: figure out how to render error here
        throw new Error('useUser must be used within a UserProvider');
    }
    return context;
}
