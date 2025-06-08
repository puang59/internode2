'use client'
import SearchHeader from "@/components/SearchHeader";
import useStore from "@/hooks/useStore";
import { useEffect, use } from "react";

export default function SearchPage({ params }: { params: Promise<{ query: string }> }) {
  const { query } = use(params);
  const setZustandQuery = useStore((s) => s.setQuery);

  useEffect(() => {
    setZustandQuery(query);
  }, [query, setZustandQuery]);

  return (
    <section>
      <SearchHeader />
      <p className="mt-50">Query: {query}</p>
    </section>
  );
}
