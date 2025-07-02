const data = [
    {
        certification: false,
        title: "Computer Science B.S.",
        institution: "University of Colorado Boulder",
        dates: "Aug 2022 - May 2026",
    },
    {
        certification: true,
        title: "Google Cybersecurity Professional Certificate",
        institution: "Coursera",
        dates: "Jun 2024",
    },
];
export default function Page() {
    return (
        <div className="flex flex-col gap-4">
            {data.map((item, index) => (
                <div
                    key={index}
                    className={`col-span-2 bg-card text-card-foreground flex flex-col gap-2  p-6  @container/card ${
                        index % 2 === 0 || index === 0
                            ? "col-start-1"
                            : "col-start-2"
                    } `}
                >
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        {item.title}
                    </h4>
                    <small className="text-sm leading-none font-medium">
                        {item.institution}, {item.dates}
                    </small>
        
                </div>
            ))}
        </div>
    );
}
