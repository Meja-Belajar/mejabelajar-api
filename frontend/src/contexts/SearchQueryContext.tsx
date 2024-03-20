import { createContext, useState } from "react";

interface Children {
  children: React.ReactNode
}

interface SearchQuery {
  query?: string;
  setQuery?: (c: string) => void;
}

const QueryContext = createContext<SearchQuery>({});

const QueryProvider = ( { children } : Children ) => {
  const [query, setQuery] = useState<string>("");
  
  return(
    <QueryContext.Provider value={{ query, setQuery }}>
      { children }
    </QueryContext.Provider>
  )
}

export {
  QueryProvider,
  QueryContext
}