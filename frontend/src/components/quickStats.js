import React, {useState, useEffect} from "react";
import { Doughnut, Pie } from 'react-chartjs-2';
import colors from '../static/colors'
import {Link} from '@reach/router'
import './quickStats.css'

const QuickStats = (props) => {
    
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

    const [topCollaborators, setTopCollaborators] = useState({loading: true})

    const getTopCollaborators = (repos) => {

        let collabCount = {}

        for (const repo of repos) {
            console.log(repo)
            if (repo.Collaborators) {
                for (const user of repo.Collaborators) {
                    if (user) {
                        if (collabCount[user]) {
                            console.log(collabCount[user])
                            collabCount[user] += 1
                        } else {
                            collabCount[user] = 1
                        }
                    }
                }
            }
        }

        let colArray = []

        for (const key in collabCount) {
            colArray.push([key, collabCount[key]])
        }

        colArray.sort((a, b) => (
            b[1] - a[1]
        ))

        colArray = colArray.slice(0,5)

        setTopCollaborators(colArray)
    }

    useEffect(() => getTopCollaborators(props.user.Repositories), [props.user.Login])

    const getCommitsPerDay = () => {
        const days = ((new Date().getTime()) - Date.parse(props.user.CreatedAt)) / (1000*60*60*24)
        return (props.user.TotalContributions)/days
    }

    return (
        <>
        <ul id="big-nums">
            <li>
                <span>{props.user.TotalContributions}</span>
                Commits
            </li>
            <li>
                <span>{props.user.PullRequestsMade}</span>
                Pull Requests
            </li>
            <li>
                <span>{getCommitsPerDay().toFixed(2)}</span>
                Commits/Day
            </li>
            <li>
                <span>{props.user.Stars}</span>
                Stars
            </li>
        </ul>
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
            <div></div>
            <div>
                <h3>Frequent Collaborators</h3>
                {
                    topCollaborators.loading 
                    ?
                        <>Loading...</>
                    :
                        <table>
                            <tbody>
                                <tr>
                                    <th>
                                        User
                                    </th>
                                    <th className="numeric">
                                        Collaborations
                                    </th>
                                </tr>
                            {
                                topCollaborators.map((col) => (
                                    <tr key={col[0] + col[1]}>
                                        <td><Link to={`/${col[0]}`}>@{col[0]}</Link></td>
                                        <td className="numeric">{col[1]}</td>
                                    </tr>
                                ))
                            }
                            </tbody>
                        </table>
                }
            </div>
        </section>
        </>
    );
}

export default QuickStats;