import { create } from 'zustand';

export type ResultType = {
  query: string;
  results: string[];
}

type BearState = {
  query: string;
  result: ResultType;
  setQuery: (query: string) => void;
  setResult: (result: ResultType) => void;
};

const useStore = create<BearState>((set) => ({
  query: '',
  result: { query: '', results: [] },
  setQuery: (query) => set({ query }),
  setResult: (result) => set({ result }),
}));

export default useStore;
