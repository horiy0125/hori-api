class ApiV1::ExternalPostsController < ApplicationController

  def index
    @external_posts = ExternalPost.order(published_at: :desc)
  end
end
