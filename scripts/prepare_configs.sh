#!/bin/sh

CONFIG_FLAG=$1
PATH_FLAG=$2
FORCE_FLAG=$4
WARNING="WRONG USAGE! please type --help for usage."

if [ "$3" = "" ]; then PATH="$(pwd)/config"; else PATH=$3; fi

handle_path()
{

	if [ "$FORCE_FLAG" = "-f" ]; then force=true; else force=false; fi
	
	if [ ! -d "$PATH" ] && [ "$force" = false ]; then
		echo "The given path does not exist. use -f after path to force."
		return
	fi	

}


handle_yaml()
{

		echo "--yaml is not active right now. Use --env instead."
		return
}

handle_env()
{
		export SERVER_HOST='localhost'
 	        export SERVER_PORT=8000
		export READ_TIMEOUT=10
		export WRITE_TIMEOUT=600	
}


handle_config_flag()
{
case $CONFIG_FLAG in
	--path)
		echo $WARNING 
		return
		;;

	--yaml)
		handle_yaml
		;;
	--env)
		handle_env
		;;
	--help)
		echo "\nOPTIONAL FLAG(S):"
		echo "			--path <PATH FOR CONFIG FILE>"
		echo "\nREQUIRED FLAG(S):"
		echo "			--yaml"
		echo "			--env"
		echo "EXAMPLE USAGE:	./${0##*/} --yaml --path /home/myproject/config"		
		return
		;;

	*)
		echo $WARNING
		return
		;;
esac
}

handle_path_flag()
{

case $PATH_FLAG in
	"")	
		handle_path
		;;
	"--path")
		handle_path
		;;
	*)
		echo $WARNING
		return
		;;
esac

}

handle_path_flag
handle_config_flag

echo "done"
