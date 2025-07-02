export default function Page() {
    return (
        <div className="w-full h-full flex flex-col gap-2 items-center justify-center">
            <h1 className="scroll-m-20 text-center text-4xl font-extrabold tracking-tight text-balance">
                Paul Rodriguez
            </h1>
            <p className="text-muted-foreground text-xl">
                Student - Engineer - Explorer - Music Seeker
            </p>
            <p className="leading-7 [&:not(:first-child)]:mt-6 max-w-xl">
                Hello and Welcome! I&rsquo;m Paul. I love to build solutions that directly solve problems in my own life. I learned to program during the first COVID summer, and quickly realized I would like to make it my career. I have co-founded a business, and while it eventually failed, I learned invaluable hard/soft skills. Afterwards I decided to go to CU Boulder for CS to further my knowledge. In my free time I love to travel to wherever the latest DJ I am bumping is playing a set, and working on personal projects.
            </p>
        </div>
    );
}
