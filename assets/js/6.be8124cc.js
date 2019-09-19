(window.webpackJsonp=window.webpackJsonp||[]).push([[6],{218:function(e,s,a){"use strict";a.r(s);var t=a(0),n=Object(t.a)({},(function(){var e=this,s=e.$createElement,a=e._self._c||s;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"starting-a-local-chain"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#starting-a-local-chain","aria-hidden":"true"}},[e._v("#")]),e._v(" Starting a local chain")]),e._v(" "),a("p",[e._v("Inside the following page you will learn how to start a new Commercio.network chain that might be useful to you\nin order to perform some tests without connecting to the testnet or mainnet.")]),e._v(" "),a("h2",{attrs:{id:"installation"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#installation","aria-hidden":"true"}},[e._v("#")]),e._v(" Installation")]),e._v(" "),a("p",[e._v("In order to start a local test chain you will need to install the latest "),a("code",[e._v("cnd")]),e._v(" and "),a("code",[e._v("cncli")]),e._v(" binaries.\nTo do so, please execute the following commands:")]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[e._v("git")]),e._v(" clone https://github.com/commercionetwork/commercionetwork\n"),a("span",{pre:!0,attrs:{class:"token builtin class-name"}},[e._v("cd")]),e._v(" commercionetwork\n"),a("span",{pre:!0,attrs:{class:"token function"}},[e._v("make")]),e._v(" "),a("span",{pre:!0,attrs:{class:"token function"}},[e._v("install")]),e._v("\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br"),a("span",{staticClass:"line-number"},[e._v("3")]),a("br")])]),a("p",[e._v("The output should look like the following:")]),e._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v('GO111MODULE=on go install -tags " ledger" ./cmd/cnd\nGO111MODULE=on go install -tags " ledger" ./cmd/cncli\n')])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br")])]),a("p",[e._v("Now, you should be able to execute the following command:")]),e._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("cnd version\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("p",[e._v("If the version number is printed properly, you are ready to go.")]),e._v(" "),a("h2",{attrs:{id:"chain-starting"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#chain-starting","aria-hidden":"true"}},[e._v("#")]),e._v(" Chain starting")]),e._v(" "),a("p",[e._v("In order to start the chain, the following steps must be performed:")]),e._v(" "),a("ol",[a("li",[a("a",{attrs:{href:"#1-resetting-previous-instances"}},[e._v("Resetting previous instances")])]),e._v(" "),a("li",[a("a",{attrs:{href:"#2-init-a-new-chain"}},[e._v("Init a new chain")])]),e._v(" "),a("li",[a("a",{attrs:{href:"#3-setup-the-genesis-data"}},[e._v("Setup the genesis data")])]),e._v(" "),a("li",[a("a",{attrs:{href:"#4-collect-the-genesis-transactions"}},[e._v("Collect the genesis transactions")])]),e._v(" "),a("li",[a("a",{attrs:{href:"#5-start-the-chain"}},[e._v("Start the chain")])])]),e._v(" "),a("h3",{attrs:{id:"_1-resetting-previous-instances"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_1-resetting-previous-instances","aria-hidden":"true"}},[e._v("#")]),e._v(" 1. Resetting previous instances")]),e._v(" "),a("p",[e._v("In order to start a chain without any problem, you will need to reset everything.\nTo do so, execute the following commands:")]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[a("span",{pre:!0,attrs:{class:"token function"}},[e._v("rm")]),e._v(" -r ~/.cnd\ncnd unsafe-reset-all\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br")])]),a("div",{staticClass:"warning custom-block"},[a("p",[e._v("This will remove all the previous chain data so please make sure to backup\nthe "),a("code",[e._v("~/.cnd")]),e._v(" folder just in case you need the data back later.")])]),e._v(" "),a("h3",{attrs:{id:"_2-init-a-new-chain"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_2-init-a-new-chain","aria-hidden":"true"}},[e._v("#")]),e._v(" 2. Init a new chain")]),e._v(" "),a("p",[e._v("To initialize a new chain, please execute the following command:")]),e._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("cnd init testchain --overwrite\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("h3",{attrs:{id:"_3-setup-the-genesis-data"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_3-setup-the-genesis-data","aria-hidden":"true"}},[e._v("#")]),e._v(" 3. Setup the genesis data")]),e._v(" "),a("p",[e._v("Now that you have initialized the new chain, you need to set some genesis data."),a("br"),e._v("\nTo do so we will use some commands that require you to have a local account key name and password.\nIf you haven't create one yet, please do it know by executing")]),e._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[e._v("cncli keys "),a("span",{pre:!0,attrs:{class:"token function"}},[e._v("add")]),e._v(" jack\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("p",[e._v("After this command, please insert a password that will be later used.")]),e._v(" "),a("div",{staticClass:"warning custom-block"},[a("p",[e._v("While creating a local key please use a password that you will remember easily as it will be used\noften later during the procedure.")])]),e._v(" "),a("p",[e._v("The output to the previous command should look something like the following:")]),e._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v('- name: jack\n  type: local\n  address: did:com:15erw8aqttln5semks0vnqjy9yzrygzmjwh7vke\n  pubkey: did:com:pub1addwnpepqgkyyqvz2e3um89luc34wt4rlhv63jlgky6eyvc4x57ee8hngl8z2h3d3zn\n  mnemonic: ""\n  threshold: 0\n  pubkeys: []\n\n\n**Important** write this mnemonic phrase in a safe place.\nIt is the only way to recover your account if you ever forget your password.\n\nmiddle vanish genre gorilla label insane east need holiday fringe odor ice permit hen art benefit amazing worry evoke pigeon project van unfold fantasy\n')])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br"),a("span",{staticClass:"line-number"},[e._v("3")]),a("br"),a("span",{staticClass:"line-number"},[e._v("4")]),a("br"),a("span",{staticClass:"line-number"},[e._v("5")]),a("br"),a("span",{staticClass:"line-number"},[e._v("6")]),a("br"),a("span",{staticClass:"line-number"},[e._v("7")]),a("br"),a("span",{staticClass:"line-number"},[e._v("8")]),a("br"),a("span",{staticClass:"line-number"},[e._v("9")]),a("br"),a("span",{staticClass:"line-number"},[e._v("10")]),a("br"),a("span",{staticClass:"line-number"},[e._v("11")]),a("br"),a("span",{staticClass:"line-number"},[e._v("12")]),a("br"),a("span",{staticClass:"line-number"},[e._v("13")]),a("br")])]),a("p",[e._v("Once you have create a local key, you can execute the following commands:")]),e._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# Add some funds to the account")]),e._v("\ncnd add-genesis-account "),a("span",{pre:!0,attrs:{class:"token variable"}},[a("span",{pre:!0,attrs:{class:"token variable"}},[e._v("$(")]),e._v("cncli keys show jack --address"),a("span",{pre:!0,attrs:{class:"token variable"}},[e._v(")")])]),e._v(" 10000000000000ucommercio\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# Set the account to be the government")]),e._v("\ncnd set-genesis-government-address "),a("span",{pre:!0,attrs:{class:"token variable"}},[a("span",{pre:!0,attrs:{class:"token variable"}},[e._v("$(")]),e._v("cncli keys show jack --address"),a("span",{pre:!0,attrs:{class:"token variable"}},[e._v(")")])]),e._v("\n\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[e._v("# Optional - Set the account to be a membership minter")]),e._v("\ncnd add-genesis-membership-minter "),a("span",{pre:!0,attrs:{class:"token variable"}},[a("span",{pre:!0,attrs:{class:"token variable"}},[e._v("$(")]),e._v("cncli keys show jack --address"),a("span",{pre:!0,attrs:{class:"token variable"}},[e._v(")")])]),e._v("\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br"),a("span",{staticClass:"line-number"},[e._v("3")]),a("br"),a("span",{staticClass:"line-number"},[e._v("4")]),a("br"),a("span",{staticClass:"line-number"},[e._v("5")]),a("br"),a("span",{staticClass:"line-number"},[e._v("6")]),a("br"),a("span",{staticClass:"line-number"},[e._v("7")]),a("br"),a("span",{staticClass:"line-number"},[e._v("8")]),a("br")])]),a("p",[e._v("After executing those commands, make sure your genesis file is valid by executing")]),e._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[e._v("cnd validate-genesis\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("p",[e._v("This should output something similar to the following text:")]),e._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("validating genesis file at /home/user/.cnd/config/genesis.json\nFile at /home/user/.cnd/config/genesis.json is a valid genesis file\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br")])]),a("h3",{attrs:{id:"_4-collect-the-genesis-transactions"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_4-collect-the-genesis-transactions","aria-hidden":"true"}},[e._v("#")]),e._v(" 4. Collect the genesis transactions")]),e._v(" "),a("p",[e._v("Once you've setup the genesis file, you can create the genesis transaction and collect it.\nTo do so, please run")]),e._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[e._v("cnd gentx --name jack --amount 100000000ucommercio\ncnd collect-gentxs\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br")])]),a("h3",{attrs:{id:"_5-start-the-chain"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#_5-start-the-chain","aria-hidden":"true"}},[e._v("#")]),e._v(" 5. Start the chain")]),e._v(" "),a("p",[e._v("Once all the genesis transactions have been created, you can start the chain by running")]),e._v(" "),a("div",{staticClass:"language-shell line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-shell"}},[a("code",[e._v("cnd start\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br")])]),a("p",[e._v("You should now be able to see an output that looks something like the following:")]),e._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("I[2019-09-19|10:26:06.651] Starting ABCI with Tendermint                module=main \nI[2019-09-19|10:26:12.034] Executed block                               module=state height=1 validTxs=0 invalidTxs=0\nI[2019-09-19|10:26:12.046] Committed state                              module=state height=1 txs=0 appHash=522AF70477C8C53361489DB2D592BF66C37E76C52A42DC7AE8230AD76EF3B54F\nI[2019-09-19|10:26:17.128] Executed block                               module=state height=2 validTxs=0 invalidTxs=0\nI[2019-09-19|10:26:17.140] Committed state                              module=state height=2 txs=0 appHash=8BD8E4D3D66A60C37B1AE721E2C7B259C36A65209575A548CB4D09BEF0B0E42E\n...\n")])]),e._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[e._v("1")]),a("br"),a("span",{staticClass:"line-number"},[e._v("2")]),a("br"),a("span",{staticClass:"line-number"},[e._v("3")]),a("br"),a("span",{staticClass:"line-number"},[e._v("4")]),a("br"),a("span",{staticClass:"line-number"},[e._v("5")]),a("br"),a("span",{staticClass:"line-number"},[e._v("6")]),a("br")])])])}),[],!1,null,null,null);s.default=n.exports}}]);