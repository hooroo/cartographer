require './web_app'

run Rack::URLMap.new(
  '/' => WebApp,
)
