'use client';
import { useRouter } from "next/navigation";
import { Input } from "./ui/input"
import useStore from "@/hooks/useStore";
import { useEffect } from "react";

export default function SearchHeader() {
  const router = useRouter();
  const query = useStore((s) => s.query);
  const result = useStore((s) => s.result);

  const setZustandResult = useStore((s) => s.setResult)

  useEffect(() => {
    const quickSearch = async () => {
      const response = await fetch(`http://localhost:8080/quicksearch?q=${query}`)
      if (!response.ok) {
        console.error('Failed to fetch search results');
        return;
      }
      const data = await response.json();
      setZustandResult(data);
      console.log(data);
    }
    if (query) {
      quickSearch();
    }
  }, [query]);

  return (
    <section className="fixed top-0 w-full z-50 bg-[#f7f7f7] px-4 py-4 flex items-center justify-between">
      <div className="flex space-x-10">
        <h1 className="text-2xl font-bold font-instrument-serif italic cursor-pointer" onClick={() => router.push('/')}>internode</h1>
        <Input
          type="text"
          placeholder="Search..."
          className="rounded-full w-98"
          onChange={(e) => console.log(e.target.value)}
        />
      </div>
    </section>
  )
}
