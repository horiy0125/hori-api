class ApiV1::QiitaPostsController < ApplicationController

  def index
    @qiita_posts = QiitaPost.order(published_at: :desc)
    render :json => @qiita_posts
  end
end
