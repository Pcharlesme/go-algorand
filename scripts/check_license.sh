#!/bin/bash

PROJECT_ROOT=$(git rev-parse --show-toplevel)
LICENSE_LOCATION="$PROJECT_ROOT"/scripts/LICENSE_HEADER
NUMLINES=$(< "$LICENSE_LOCATION" wc -l | tr -d ' ')
CURRENT_YEAR=$(date +"%Y")
LICENSE=$(sed "s/{DATE_Y}/$CURRENT_YEAR/" "$LICENSE_LOCATION")
VERSIONED_GO_FILES=$(git ls-tree --full-tree --name-only -r HEAD | grep "\.go$")
EXTRA_FILES=(
    cmd/tealdbg/bundle_home_html.sh
    crypto/memcpy_chk_windows.c
    tools/x-repo-types/typeAnalyzer/main.tmpl
    test/heapwatch/block_history.py
    test/heapwatch/block_history_plot.py
    test/heapwatch/metrics_delta.py
    test/heapwatch/nodeHostTarget.py
    test/heapwatch/client_ram_report.py
    test/heapwatch/runNodeHost.py
    test/heapwatch/block_history_relays.py
    test/heapwatch/heapWatch.py
)
EXCLUDE=(
    "Code generated by"
    "David Lazar"
    "Go Authors"
    "Google Inc"
    "Prometheus Authors"
    "Jeffrey Wilcke"
    "dummy.go"
)
FILTER=$(IFS="|" ; echo "${EXCLUDE[*]}")
INPLACE=false
UPDATE=false
VERBOSE=false
MOD_COUNT=0
RETURN_VALUE=0

usage() {
    echo "check_license"
    echo
    echo "Usage: $0 [args]"
    echo
    echo "Args:"
    echo "-i    Edit in-place."
    echo "-u    Update license to current year."
    echo "-v    Verbose, same as doing \`head -n ${NUMLINES:-15}\` on each file."
    echo
}

RED_FG=$(tput setaf 1 2>/dev/null)
END_FG_COLOR=$(tput sgr0 2>/dev/null)

while [ "$1" != "" ]; do
    case "$1" in
        -i)
            INPLACE=true
            ;;
        -u)
            UPDATE=true
            ;;
        -v) VERBOSE=true
            ;;
        -h)
            usage
            exit 0
            ;;
        *)
            echo "${RED_FG}[ERROR]${END_FG_COLOR} Unknown option $1"
            usage
            exit 1
            ;;
    esac
    shift
done

for FILE in $VERSIONED_GO_FILES; do
    # https://en.wikipedia.org/wiki/Cat_(Unix)#Useless_use_of_cat
    if [[ $LICENSE != $(<"$PROJECT_ROOT/$FILE" head -n "$NUMLINES") ]]; then

        if <"$PROJECT_ROOT/$FILE" head -n "$NUMLINES" | tr "\n" " " | grep -qvE "$FILTER"; then
            RETURN_VALUE=1

            if ! $VERBOSE; then
                if $UPDATE; then
                    sed -i.orig s/Copyright\ \(C\)\ 2019-....\ Algorand,\ Inc\./Copyright\ \(C\)\ 2019-$CURRENT_YEAR\ Algorand,\ Inc./ "$PROJECT_ROOT/$FILE" && \
                        rm "$PROJECT_ROOT/$FILE".orig
                    ((MOD_COUNT++))
                elif $INPLACE; then
                    cat <(echo -e "$LICENSE\n") "$PROJECT_ROOT/$FILE" > "$PROJECT_ROOT/$FILE".1 &&
                        mv "$PROJECT_ROOT/$FILE"{.1,}
                    ((MOD_COUNT++))
                fi
                echo "$FILE"
            else
                echo -e "\n${RED_FG}$FILE${END_FG_COLOR}"
                <"$PROJECT_ROOT/$FILE" head -n "$NUMLINES"
                echo
            fi
        fi
    fi
done

# non-go files that include a license header
for FILE in "${EXTRA_FILES[@]}"; do
    if ! grep -qs "Copyright (C) 2019-$CURRENT_YEAR Algorand, Inc." "$PROJECT_ROOT/$FILE"; then
        RETURN_VALUE=1
        if ! $VERBOSE; then
            if $UPDATE; then
                sed -i.orig s/Copyright\ \(C\)\ 2019-....\ Algorand,\ Inc\./Copyright\ \(C\)\ 2019-$CURRENT_YEAR\ Algorand,\ Inc./ "$PROJECT_ROOT/$FILE" && \
                    rm "$PROJECT_ROOT/$FILE".orig
                ((MOD_COUNT++))
            fi
            # It's dangerous to do inplace updates of non-go files, because their format might be different
            echo "$FILE"
        else
            echo -e "\n${RED_FG}$FILE${END_FG_COLOR}"
            <"$PROJECT_ROOT/$FILE" head -n "$NUMLINES"
            echo
        fi
    fi
done

# check the README.md file.
READMECOPYRIGHT="Copyright (C) 2019-$CURRENT_YEAR, Algorand Inc."
if [ "$(<README.md grep -c "${READMECOPYRIGHT}" | tr -d ' ')" = "0" ]; then
    RETURN_VALUE=1
    if ! $VERBOSE; then
        if $UPDATE; then
            sed -i.orig s/Copyright\ \(C\)\ 2019-....,\ Algorand\ Inc\./Copyright\ \(C\)\ 2019-$CURRENT_YEAR,\ Algorand\ Inc./ README.md &&
                rm README.md.orig
            ((MOD_COUNT++))
        fi
        echo "README.md"
    else
        echo -e "\n${RED_FG}README.md${END_FG_COLOR}"
        grep 'Copyright (C) 2019' README.md
        echo
    fi
fi

if [ $RETURN_VALUE -ne 0 ]; then
    echo -e "\n${RED_FG}FAILED LICENSE CHECK.${END_FG_COLOR}"
    if [ $INPLACE == "false" ] && [ $UPDATE == "false" ]; then
        echo -e "Use 'check_license -i' to install to new files, 'check_license.sh -u' to update year."
    else
        echo "Modified $MOD_COUNT file(s)."
    fi
    echo ""
fi

exit $RETURN_VALUE

