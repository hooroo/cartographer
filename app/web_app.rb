require 'sinatra/base'
require 'sinatra/reloader'
require 'better_errors'
require 'json'

require_relative './config/initialise'
require_relative './presenters/report_presenter'

module Cartographer
  class WebApp < Sinatra::Base

    configure :development do
      register Sinatra::Reloader

      APP_FILES.each do |file|
        also_reload file
      end

      use BetterErrors::Middleware
    end

    # set :views, './app/views'

    get '/' do
      @reports = get_reports

      erb :index
    end

    private

      # TODO: Make this real!
      #
      def get_reports
        reports = JSON.parse('[
          {
            "system": { "node_name": "node.name", "ip_address": "192.168.0.1" },
            "reports": {
              "unicorn": { "workers": "4", "memory_used": "4096" },
              "nginx": { "workers": "8", "memory_used": "8192" }
            }
          }
        ]')

        reports.map { |x| Presenters::ReportPresenter.new(x) }
      end

  end
end
