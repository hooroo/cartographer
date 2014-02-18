require 'hashie'

require_relative './report_presenter'

module Cartographer
  module Presenters
    class NodePresenter

      def initialize json
        @data = Hashie::Mash.new(json)
      end

      def node_name
        data.system.node_name
      end

      def reports
        data.reports.map { |identifier, json| Presenters::ReportPresenter.new(identifier, json) }
      end

      def css_status
        # warning, danger
        'success'
      end

      private

        attr_reader :data

    end
  end
end
