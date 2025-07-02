import { Button } from "@/components/ui/button";
import Link from "next/link";
export default function Page() {
    return (
        <div className="w-full h-full flex flex-col gap-4 items-center justify-center">
            <h1 className="scroll-m-20 text-center text-4xl font-extrabold tracking-tight text-balance">
                Hi, I'm Paul.
            </h1>

            <Button asChild>
                 <Link href="/bio">Explore</Link>
            </Button>
        </div>
    );
}
