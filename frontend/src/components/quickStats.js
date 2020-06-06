import React, {useState, useEffect} from "react";
import { Doughnut, Pie } from 'react-chartjs-2';
import colors from '../static/colors'
import './quickStats.css'
import {Link} from '@reach/router'

const QuickStats = (props) => {

    const [mostLines, setMostLines] = useState({loading: true})

    const getHighestLineCount = (repos) => {
        let languageCounts = {}
        for (const repo of repos) {
            for (const lang in repo.Languages) {
                if (languageCounts[lang]) {
                    languageCounts[lang] += repo.Languages[lang]
                } else {
                    languageCounts[lang] = repo.Languages[lang]
                }
            }
        }

        var sorted = [];
        for (const lang in languageCounts) {
            sorted.push([lang, languageCounts[lang]]);
        }

        sorted.sort(function(a, b) {
            return b[1] - a[1];
        });

        sorted = sorted.slice(0,10)

        let lines = [], names = [], backgrounds = []
        for (const tuple of sorted) {
            names.push(tuple[0])
            lines.push(tuple[1])
            if (colors[tuple[0]]) {
                backgrounds.push(colors[tuple[0]].color)
            } else backgrounds.push("#333333")
        }

        const data = {
            labels: names,
            datasets: [{
                data: lines,
                backgroundColor: backgrounds,
                hoverBackgroundColor: backgrounds,
            }]
        };

        setMostLines(data);
    }

    useEffect(() => getHighestLineCount(props.user.Repositories), [props.user.Login])

    const [mainLanguages, setMainLanguages] = useState({loading: true})

    const getMainLanguages = (repos) => {
        let languageCounts = {}
        for (const repo of repos) {
            if (languageCounts[repo.PrimaryLanguage]) {
                languageCounts[repo.PrimaryLanguage] += 1
            } else {
                languageCounts[repo.PrimaryLanguage] = 1
            }
        }
        let sorted = [];
        for (const lang in languageCounts) {
            if (lang) {
                sorted.push([lang, languageCounts[lang]]);
            }
        }

        sorted.sort(function(a, b) {
            return b[1] - a[1];
        });

        // sorted = sorted.slice(0, 8)

        let lines = [], names = [], backgrounds = []
        for (const tuple of sorted) {
            names.push(tuple[0])
            lines.push(tuple[1])
            if (colors[tuple[0]]) {
                backgrounds.push(colors[tuple[0]].color)
            } else backgrounds.push("#333333")
        }

        const data = {
            labels: names,
            datasets: [{
                data: lines,
                backgroundColor: backgrounds,
                hoverBackgroundColor: backgrounds,
                borderWidth: 1.5
            }],
        };

        setMainLanguages(data);
    }

    useEffect(() => getMainLanguages(props.user.Repositories), [props.user.Login])

    return (
        <section id="stats">
            <div>
                <h3>Repos per Language</h3>
                {
                    mainLanguages.loading 
                    ?
                        <>Loading...</>
                    :
                        <Doughnut 
                            data={mainLanguages}
                            legend={{position: 'left', labels: {boxWidth: 12, fontSize: 12, fontColor: "white"}}} 
                            options={{cutoutPercentage: 40}}
                        />
                }
            </div>
            <div>
            {/* {
                mostLines.loading 
                ?
                    <>Loading...</>
                :
                    <Doughnut data={mostLines} legend={{position: 'left'}} />
            } */}

            {/* <ul>
                        {
                            mostLines.slice(0,5).map(lang => (
                                <li key={lang}>{lang[0]}</li>
                            ))
                        }
                    </ul> */}
            </div>
            <div></div>
        </section>
    );
}

export default QuickStats;