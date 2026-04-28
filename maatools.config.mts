import type {FullConfig} from "@nekosu/maa-tools";

const config: FullConfig = {
    cwd: import.meta.dirname,
    maaVersion: "latest",
    interfacePath: "assets/interface.json",
    check: {},
    vscode: {
        agents: {
            "agent/go-service": "launch-go-agent",
        },
    },
};

export default config;
