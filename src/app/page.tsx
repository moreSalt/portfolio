import { Button } from "@/components/ui/button";
import { FrownIcon} from "lucide-react";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";

import Link from "next/link";
export default function Page() {
    return (
        <div className="w-full h-full flex flex-col gap-4 items-center justify-center">
            <h1 className="scroll-m-20 text-center text-4xl font-extrabold tracking-tight text-balance">
                Hi, I&rsquo;m Paul.
            </h1>
            <Alert className="w-full max-w-lg">
                <FrownIcon />
                <AlertTitle>I&rsquo;m ruined</AlertTitle>
                <AlertDescription>
                    Giving your credit card to open ai is a dangerous game
                </AlertDescription>
            </Alert>
            <Button asChild>
                <Link href="/bio">Explore</Link>
            </Button>
        </div>
    );
}
