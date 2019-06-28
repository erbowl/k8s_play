while sleep 0.01; do curl -s -w " -- `date` -- %{http_code}\n" $DEMO_URL; done

