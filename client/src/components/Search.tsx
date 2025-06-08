'use client'
import { Input } from "./ui/input"
import { Button } from "./ui/button"
import { useState, useRef } from "react"
import { useRouter } from "next/navigation"

export default function Search() {
  const router = useRouter();
  const buttonRef = useRef<HTMLButtonElement>(null);
  const [searchQuery, setSearchQuery] = useState<string>();

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      handleRedirect();
    }
  }

  const handleRedirect = async () => {
    // const response = await fetch(`http://localhost:8080/quicksearch?q=${encodeURIComponent(searchQuery || '')}`)
    // const data = await response.json();
    // console.log(data);
    router.push(`/search/${encodeURIComponent(searchQuery || '')}`);
  }

  return (
    <section>
      <div className="flex items-center justify-center w-full h-16 ">
        <Input
          type="text"
          placeholder="Search..."
          onChange={(e) => setSearchQuery(e.target.value)}
          onKeyDown={handleKeyDown}
        />
        <Button
          className="ml-2"
          onClick={handleRedirect}
          ref={buttonRef}
        >
          Search
        </Button>
      </div>
    </section>
  )
}
