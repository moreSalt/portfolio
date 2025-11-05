const data = [
    {
        title: "Travel Project",
        hook: "Find your next travel destination based on your favorite musical artists",
        description:
            "I was trying to figure out where I could go during spring break to see some of my favorite artists. I built this using Golang for the backend, and then the front-end was originally built in Sveltekit but then switched to Next. Hardest part of this project was getting wait times down from 3+ mins to under 30s.",
        url: "https://travel.fluster.dev",
    },
    {
        title: "Speech to Latex",
        hook: "It would be nice if my math professor had lecture notes written in Latex",
        description:
            "This was made by myself and my team for my favorite class 'Big Data Architecture'. Using AI we built an app that takes in speech and converts it into Latex. ",
        url: "https://github.com/moreSalt/speechtolatex-client",
    },
    {
        title: "HelloFresh to Instacart",
        hook: "I hate finding recipes, I hate grocery shopping",
        description:
            "I built this script to simplify my life. Basically it will take any HelloFresh recipe and then return all of the ingredients required in Instacart. There was also a version of this that hooked up directly to Kroger Grocery.",
        url: "https://github.com/moreSalt/hellofresh-instacart",
    },
    {
        title: "Jobs Data",
        hook: "Searching for trends from job posts",
        description: "This is my current project. I was curious about a few things. 1. What insights can be found from looking at a large data set of job postings. 2. Why do some companies slightly update job postings every day? 3. Does finding a job posting sooner result in more interviews.",
        url: "#"
    },
    {
        title: "More?",
        hook: "Want to hear about what else I have worked on?",
        description: "There is a whole bunch of stuff in the private repos, but in varying states of done, broken, pivoted, etc. Reach out to me via email to learn more.",
        url: "https://fluster.dev"
    },
];

export default function Page() {
    return (
        <div className="grid grid-cols-3 gap-4">
            {data.map((item, index) => (
                <div className="aspect-square bg-card text-card-foreground flex flex-col gap-2 p-6 @container/card" key={index}>
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        <a href={item.url} target="_blank">
                            {item.title}
                        </a>
                    </h4>
                    <p className="text-muted-foreground text-xl">
                        {item.hook}
                    </p>
                    <p className="leading-7">{item.description}</p>
                </div>
            ))}
        </div>
    );
}
