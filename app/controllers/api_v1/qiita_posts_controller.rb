class ApiV1::QiitaPostsController < ApplicationController

  def index
    @qiita_posts = QiitaPost.order(published_at: :desc)
  end
end
