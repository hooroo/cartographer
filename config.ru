require './app/web_app'

run Rack::URLMap.new(
  '/' => Cartographer::WebApp,
)
